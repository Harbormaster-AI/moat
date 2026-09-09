import React, { Component } from 'react'
import InvestmentPortfolioService from '../services/InvestmentPortfolioService'

class ListInvestmentPortfolioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                investmentPortfolios: []
        }
        this.addInvestmentPortfolio = this.addInvestmentPortfolio.bind(this);
        this.editInvestmentPortfolio = this.editInvestmentPortfolio.bind(this);
        this.deleteInvestmentPortfolio = this.deleteInvestmentPortfolio.bind(this);
    }

    deleteInvestmentPortfolio(id){
        InvestmentPortfolioService.deleteInvestmentPortfolio(id).then( res => {
            this.setState({investmentPortfolios: this.state.investmentPortfolios.filter(investmentPortfolio => investmentPortfolio.investmentPortfolioId !== id)});
        });
    }
    viewInvestmentPortfolio(id){
        this.props.history.push(`/view-investmentPortfolio/${id}`);
    }
    editInvestmentPortfolio(id){
        this.props.history.push(`/add-investmentPortfolio/${id}`);
    }

    componentDidMount(){
        InvestmentPortfolioService.getInvestmentPortfolios().then((res) => {
            this.setState({ investmentPortfolios: res.data});
        });
    }

    addInvestmentPortfolio(){
        this.props.history.push('/add-investmentPortfolio/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InvestmentPortfolio List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInvestmentPortfolio}> Add InvestmentPortfolio</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PortfolioCode </th>
                                    <th> BaseCurrency </th>
                                    <th> CreatedAt </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.investmentPortfolios.map(
                                        investmentPortfolio => 
                                        <tr key = {investmentPortfolio.investmentPortfolioId}>
                                             <td> { investmentPortfolio.portfolioCode } </td>
                                             <td> { investmentPortfolio.baseCurrency } </td>
                                             <td> { investmentPortfolio.createdAt } </td>
                                             <td> { investmentPortfolio.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInvestmentPortfolio(investmentPortfolio.investmentPortfolioId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInvestmentPortfolio(investmentPortfolio.investmentPortfolioId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInvestmentPortfolio(investmentPortfolio.investmentPortfolioId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListInvestmentPortfolioComponent
