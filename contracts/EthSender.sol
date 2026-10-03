// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @title EthSender
/// @notice A faucet: anyone may claim `amountPerClaim` once per `cooldown`.
/// The earlier version let any caller claim 42 ETH again and again until the
/// contract was empty.
contract EthSender {
    address public immutable owner;
    uint256 public immutable amountPerClaim;
    uint256 public immutable cooldown;

    /// @notice Total ETH sent to each address.
    mapping(address => uint256) public totalSent;
    /// @notice When each address last claimed.
    mapping(address => uint256) public lastClaim;

    event EthSent(address indexed recipient, uint256 amount);
    event Withdrawn(address indexed to, uint256 amount);

    error NotOwner();
    error TooSoon(uint256 nextClaimAt);
    error InsufficientBalance(uint256 available, uint256 required);
    error TransferFailed();

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    constructor(uint256 amountPerClaim_, uint256 cooldown_) {
        owner = msg.sender;
        amountPerClaim = amountPerClaim_;
        cooldown = cooldown_;
    }

    /// @notice Sends `amountPerClaim` to the caller.
    function sendEth() external {
        uint256 last = lastClaim[msg.sender];
        if (last != 0 && block.timestamp < last + cooldown) revert TooSoon(last + cooldown);
        if (address(this).balance < amountPerClaim) {
            revert InsufficientBalance(address(this).balance, amountPerClaim);
        }

        // Effects before the external call, so a re-entering caller sees its claim.
        lastClaim[msg.sender] = block.timestamp;
        totalSent[msg.sender] += amountPerClaim;
        emit EthSent(msg.sender, amountPerClaim);

        (bool ok,) = payable(msg.sender).call{value: amountPerClaim}("");
        if (!ok) revert TransferFailed();
    }

    function getTotalSent(address account) external view returns (uint256) {
        return totalSent[account];
    }

    receive() external payable {}

    /// @notice Lets the owner take ETH back out of the faucet.
    function withdraw(uint256 amount) external onlyOwner {
        if (address(this).balance < amount) revert InsufficientBalance(address(this).balance, amount);
        emit Withdrawn(owner, amount);
        (bool ok,) = payable(owner).call{value: amount}("");
        if (!ok) revert TransferFailed();
    }
}
