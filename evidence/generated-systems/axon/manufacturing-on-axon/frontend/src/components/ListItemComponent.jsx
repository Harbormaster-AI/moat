import React, { Component } from 'react'
import ItemService from '../services/ItemService'

class ListItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                items: []
        }
        this.addItem = this.addItem.bind(this);
        this.editItem = this.editItem.bind(this);
        this.deleteItem = this.deleteItem.bind(this);
    }

    deleteItem(id){
        ItemService.deleteItem(id).then( res => {
            this.setState({items: this.state.items.filter(item => item.itemId !== id)});
        });
    }
    viewItem(id){
        this.props.history.push(`/view-item/${id}`);
    }
    editItem(id){
        this.props.history.push(`/add-item/${id}`);
    }

    componentDidMount(){
        ItemService.getItems().then((res) => {
            this.setState({ items: res.data});
        });
    }

    addItem(){
        this.props.history.push('/add-item/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Item List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addItem}> Add Item</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ItemNumber </th>
                                    <th> Name </th>
                                    <th> StandardCost </th>
                                    <th> Weight </th>
                                    <th> AsSerialControlled </th>
                                    <th> ItemType </th>
                                    <th> ProcurementType </th>
                                    <th> UnitOfMeasure </th>
                                    <th> LifecycleStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.items.map(
                                        item => 
                                        <tr key = {item.itemId}>
                                             <td> { item.itemNumber } </td>
                                             <td> { item.name } </td>
                                             <td> { item.standardCost } </td>
                                             <td> { item.weight } </td>
                                             <td> { item.asSerialControlled } </td>
                                             <td> { item.itemType } </td>
                                             <td> { item.procurementType } </td>
                                             <td> { item.unitOfMeasure } </td>
                                             <td> { item.lifecycleStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editItem(item.itemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteItem(item.itemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewItem(item.itemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListItemComponent
