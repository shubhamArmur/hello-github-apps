// this file is a WETH.sol file
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract WETH is ERC20 {aa
    event Deposit(address indexed from, uint256 amount);
    event Withdraw(address indexed to, uint256 amount);

    constructor() ERC20("Wrapped ETH", "WETH") {}

    function decimals() public pure override returns (uint8) {
        return 18;
    }

    //              send ETH
    //                 |
    //         is msg.data empty?
    //            /        \
    //           yes        no
    //          /             \
    //  receive() exists?    fallback()
    //      /  \
    //    yes   no
    //    /       \
    // receive()  fallback()

    receive() external payable {
        deposit();
    }

    fallback() external payable {
        deposit();
    }

    function deposit() public payable {
        _mint(msg.sender, msg.value);

        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) external {
        _burn(msg.sender, amount);

        emit Withdraw(msg.sender, amount);

        payable(msg.sender).transfer(amount);
    }
}
