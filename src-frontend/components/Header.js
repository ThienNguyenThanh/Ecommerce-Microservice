import logo from '../assets/amazon_logo.png';
import './Header.css'
import SearchIcon from "@mui/icons-material/Search";
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';

export default function Header(){
    return (
        <div className="header">
            <img
                className="header_logo"
                src={logo}
                alt="logo"
            />
        

        <div className='header_search'>
            <input 
                className='header_searchInput'
                type="text"
            />
            <SearchIcon className="header_searchIcon"/>
        </div>

        <div className='header_nav'>
            <div className='header_options'>
                <span className='header_optionOne'>
                    Hello guest
                </span>
                <span className='header_optionTwo'>
                    Sign in
                </span>
            </div>

            <div className='header_options'>
                <span className='header_optionOne'>
                    Return
                </span>
                <span className='header_optionTwo'>
                    Order
                </span>
            </div>

            <div className='header_options'>
                <span className='header_optionOne'>
                    Your
                </span>
                <span className='header_optionTwo'>
                    Prime
                </span>
            </div>

            <div className='header_optionBasket'>
                <ShoppingCartIcon />
                <span className='header_optionLineTwo
                header_basketCount'>0</span>
            </div>
        </div>
        </div>
    )
}