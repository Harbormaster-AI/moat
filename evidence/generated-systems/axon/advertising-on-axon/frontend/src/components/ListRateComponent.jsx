import React, { Component } from 'react'
import RateService from '../services/RateService'

class ListRateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                rates: []
        }
        this.addRate = this.addRate.bind(this);
        this.editRate = this.editRate.bind(this);
        this.deleteRate = this.deleteRate.bind(this);
    }

    deleteRate(id){
        RateService.deleteRate(id).then( res => {
            this.setState({rates: this.state.rates.filter(rate => rate.rateId !== id)});
        });
    }
    viewRate(id){
        this.props.history.push(`/view-rate/${id}`);
    }
    editRate(id){
        this.props.history.push(`/add-rate/${id}`);
    }

    componentDidMount(){
        RateService.getRates().then((res) => {
            this.setState({ rates: res.data});
        });
    }

    addRate(){
        this.props.history.push('/add-rate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Rate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRate}> Add Rate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> UnitPrice </th>
                                    <th> AdFormat </th>
                                    <th> PricingModel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.rates.map(
                                        rate => 
                                        <tr key = {rate.rateId}>
                                             <td> { rate.unitPrice } </td>
                                             <td> { rate.adFormat } </td>
                                             <td> { rate.pricingModel } </td>
                                             <td>
                                                 <button onClick={ () => this.editRate(rate.rateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRate(rate.rateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRate(rate.rateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRateComponent
