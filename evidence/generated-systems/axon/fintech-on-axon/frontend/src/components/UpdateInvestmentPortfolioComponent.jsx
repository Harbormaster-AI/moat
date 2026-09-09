import React, { Component } from 'react'
import InvestmentPortfolioService from '../services/InvestmentPortfolioService';

class UpdateInvestmentPortfolioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                portfolioCode: '',
                baseCurrency: '',
                createdAt: '',
                status: ''
        }
        this.updateInvestmentPortfolio = this.updateInvestmentPortfolio.bind(this);

        this.changeportfolioCodeHandler = this.changeportfolioCodeHandler.bind(this);
        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InvestmentPortfolioService.getInvestmentPortfolioById(this.state.id).then( (res) =>{
            let investmentPortfolio = res.data;
            this.setState({
                portfolioCode: investmentPortfolio.portfolioCode,
                baseCurrency: investmentPortfolio.baseCurrency,
                createdAt: investmentPortfolio.createdAt,
                status: investmentPortfolio.status
            });
        });
    }

    updateInvestmentPortfolio = (e) => {
        e.preventDefault();
        let investmentPortfolio = {
            investmentPortfolioId: this.state.id,
            portfolioCode: this.state.portfolioCode,
            baseCurrency: this.state.baseCurrency,
            createdAt: this.state.createdAt,
            status: this.state.status
        };
        console.log('investmentPortfolio => ' + JSON.stringify(investmentPortfolio));
        console.log('id => ' + JSON.stringify(this.state.id));
        InvestmentPortfolioService.updateInvestmentPortfolio(investmentPortfolio).then( res => {
            this.props.history.push('/investmentPortfolios');
        });
    }

    changeportfolioCodeHandler= (event) => {
        this.setState({portfolioCode: event.target.value});
    }
    changebaseCurrencyHandler= (event) => {
        this.setState({baseCurrency: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/investmentPortfolios');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InvestmentPortfolio</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> portfolioCode: </label>
                                                <input placeholder="portfolioCode" name="portfolioCode" className="form-control" value={this.state.portfolioCode} onChange={this.changeportfolioCodeHandler}/>

                                            <label> baseCurrency: </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> createdAt: </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInvestmentPortfolio}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateInvestmentPortfolioComponent
