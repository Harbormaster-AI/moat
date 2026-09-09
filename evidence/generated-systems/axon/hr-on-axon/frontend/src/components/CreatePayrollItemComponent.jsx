import React, { Component } from 'react'
import PayrollItemService from '../services/PayrollItemService';

class CreatePayrollItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                amount: '',
                taxable: '',
                itemType: ''
        }
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changetaxableHandler = this.changetaxableHandler.bind(this);
        this.changeItemTypeHandler = this.changeItemTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PayrollItemService.getPayrollItemById(this.state.id).then( (res) =>{
                let payrollItem = res.data;
                this.setState({
                    amount: payrollItem.amount,
                    taxable: payrollItem.taxable,
                    itemType: payrollItem.itemType
                });
            });
        }        
    }
    saveOrUpdatePayrollItem = (e) => {
        e.preventDefault();
        let payrollItem = {
                payrollItemId: this.state.id,
                amount: this.state.amount,
                taxable: this.state.taxable,
                itemType: this.state.itemType
            };
        console.log('payrollItem => ' + JSON.stringify(payrollItem));

        // step 5
        if(this.state.id === '_add'){
            payrollItem.payrollItemId=''
            PayrollItemService.createPayrollItem(payrollItem).then(res =>{
                this.props.history.push('/payrollItems');
            });
        }else{
            PayrollItemService.updatePayrollItem(payrollItem).then( res => {
                this.props.history.push('/payrollItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PayrollItem</h3>
        }else{
            return <h3 className="text-center">Update PayrollItem</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> taxable:&emsp; </label>
                                                <input type="checkbox" placeholder="taxable" name="taxable" className="form-control" value={this.state.taxable} onChange={this.changetaxableHandler}/>


                                            <label> ItemType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePayrollItem}>Save</button>
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

export default CreatePayrollItemComponent
