import React, { Component } from 'react'
import PayrollItemService from '../services/PayrollItemService';

class UpdatePayrollItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                amount: '',
                taxable: '',
                itemType: ''
        }
        this.updatePayrollItem = this.updatePayrollItem.bind(this);

        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changetaxableHandler = this.changetaxableHandler.bind(this);
        this.changeItemTypeHandler = this.changeItemTypeHandler.bind(this);
    }

    componentDidMount(){
        PayrollItemService.getPayrollItemById(this.state.id).then( (res) =>{
            let payrollItem = res.data;
            this.setState({
                amount: payrollItem.amount,
                taxable: payrollItem.taxable,
                itemType: payrollItem.itemType
            });
        });
    }

    updatePayrollItem = (e) => {
        e.preventDefault();
        let payrollItem = {
            payrollItemId: this.state.id,
            amount: this.state.amount,
            taxable: this.state.taxable,
            itemType: this.state.itemType
        };
        console.log('payrollItem => ' + JSON.stringify(payrollItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        PayrollItemService.updatePayrollItem(payrollItem).then( res => {
            this.props.history.push('/payrollItems');
        });
    }

    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changetaxableHandler= (event) => {
        this.setState({taxable: event.target.value});
    }
    changeItemTypeHandler= (event) => {
        this.setState({itemType: event.target.value});
    }

    cancel(){
        this.props.history.push('/payrollItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PayrollItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> taxable: </label>
                                                <input type="checkbox" placeholder="taxable" name="taxable" className="form-control" value={this.state.taxable} onChange={this.changetaxableHandler}/>


                                            <label> ItemType: </label>
                                                <select value={this.state.itemType} onChange={this.changeItemTypeHandler}>
                      <option name="ItemType" className="form-control" >
                          Earning
                      </option>
                      <option name="ItemType" className="form-control" >
                          Deduction
                      </option>
                      <option name="ItemType" className="form-control" >
                          Tax
                      </option>
                      <option name="ItemType" className="form-control" >
                          Benefit
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePayrollItem}>Save</button>
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

export default UpdatePayrollItemComponent
