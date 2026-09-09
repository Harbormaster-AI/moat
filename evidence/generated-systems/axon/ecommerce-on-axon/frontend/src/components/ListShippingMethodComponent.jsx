import React, { Component } from 'react'
import ShippingMethodService from '../services/ShippingMethodService'

class ListShippingMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                shippingMethods: []
        }
        this.addShippingMethod = this.addShippingMethod.bind(this);
        this.editShippingMethod = this.editShippingMethod.bind(this);
        this.deleteShippingMethod = this.deleteShippingMethod.bind(this);
    }

    deleteShippingMethod(id){
        ShippingMethodService.deleteShippingMethod(id).then( res => {
            this.setState({shippingMethods: this.state.shippingMethods.filter(shippingMethod => shippingMethod.shippingMethodId !== id)});
        });
    }
    viewShippingMethod(id){
        this.props.history.push(`/view-shippingMethod/${id}`);
    }
    editShippingMethod(id){
        this.props.history.push(`/add-shippingMethod/${id}`);
    }

    componentDidMount(){
        ShippingMethodService.getShippingMethods().then((res) => {
            this.setState({ shippingMethods: res.data});
        });
    }

    addShippingMethod(){
        this.props.history.push('/add-shippingMethod/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ShippingMethod List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addShippingMethod}> Add ShippingMethod</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> FlatRate </th>
                                    <th> EstimatedDays </th>
                                    <th> AsActive </th>
                                    <th> MethodType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.shippingMethods.map(
                                        shippingMethod => 
                                        <tr key = {shippingMethod.shippingMethodId}>
                                             <td> { shippingMethod.name } </td>
                                             <td> { shippingMethod.flatRate } </td>
                                             <td> { shippingMethod.estimatedDays } </td>
                                             <td> { shippingMethod.asActive } </td>
                                             <td> { shippingMethod.methodType } </td>
                                             <td>
                                                 <button onClick={ () => this.editShippingMethod(shippingMethod.shippingMethodId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteShippingMethod(shippingMethod.shippingMethodId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewShippingMethod(shippingMethod.shippingMethodId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListShippingMethodComponent
