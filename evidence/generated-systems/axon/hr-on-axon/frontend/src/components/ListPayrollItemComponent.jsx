import React, { Component } from 'react'
import PayrollItemService from '../services/PayrollItemService'

class ListPayrollItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                payrollItems: []
        }
        this.addPayrollItem = this.addPayrollItem.bind(this);
        this.editPayrollItem = this.editPayrollItem.bind(this);
        this.deletePayrollItem = this.deletePayrollItem.bind(this);
    }

    deletePayrollItem(id){
        PayrollItemService.deletePayrollItem(id).then( res => {
            this.setState({payrollItems: this.state.payrollItems.filter(payrollItem => payrollItem.payrollItemId !== id)});
        });
    }
    viewPayrollItem(id){
        this.props.history.push(`/view-payrollItem/${id}`);
    }
    editPayrollItem(id){
        this.props.history.push(`/add-payrollItem/${id}`);
    }

    componentDidMount(){
        PayrollItemService.getPayrollItems().then((res) => {
            this.setState({ payrollItems: res.data});
        });
    }

    addPayrollItem(){
        this.props.history.push('/add-payrollItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PayrollItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPayrollItem}> Add PayrollItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Amount </th>
                                    <th> Taxable </th>
                                    <th> ItemType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.payrollItems.map(
                                        payrollItem => 
                                        <tr key = {payrollItem.payrollItemId}>
                                             <td> { payrollItem.amount } </td>
                                             <td> { payrollItem.taxable } </td>
                                             <td> { payrollItem.itemType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPayrollItem(payrollItem.payrollItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePayrollItem(payrollItem.payrollItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPayrollItem(payrollItem.payrollItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPayrollItemComponent
