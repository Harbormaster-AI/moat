import React, { Component } from 'react'
import ReturnItemService from '../services/ReturnItemService'

class ListReturnItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                returnItems: []
        }
        this.addReturnItem = this.addReturnItem.bind(this);
        this.editReturnItem = this.editReturnItem.bind(this);
        this.deleteReturnItem = this.deleteReturnItem.bind(this);
    }

    deleteReturnItem(id){
        ReturnItemService.deleteReturnItem(id).then( res => {
            this.setState({returnItems: this.state.returnItems.filter(returnItem => returnItem.returnItemId !== id)});
        });
    }
    viewReturnItem(id){
        this.props.history.push(`/view-returnItem/${id}`);
    }
    editReturnItem(id){
        this.props.history.push(`/add-returnItem/${id}`);
    }

    componentDidMount(){
        ReturnItemService.getReturnItems().then((res) => {
            this.setState({ returnItems: res.data});
        });
    }

    addReturnItem(){
        this.props.history.push('/add-returnItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ReturnItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReturnItem}> Add ReturnItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> Reason </th>
                                    <th> Condition </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.returnItems.map(
                                        returnItem => 
                                        <tr key = {returnItem.returnItemId}>
                                             <td> { returnItem.quantity } </td>
                                             <td> { returnItem.reason } </td>
                                             <td> { returnItem.condition } </td>
                                             <td>
                                                 <button onClick={ () => this.editReturnItem(returnItem.returnItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReturnItem(returnItem.returnItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReturnItem(returnItem.returnItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListReturnItemComponent
