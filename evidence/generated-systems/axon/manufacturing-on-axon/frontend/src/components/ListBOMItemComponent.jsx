import React, { Component } from 'react'
import BOMItemService from '../services/BOMItemService'

class ListBOMItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                bOMItems: []
        }
        this.addBOMItem = this.addBOMItem.bind(this);
        this.editBOMItem = this.editBOMItem.bind(this);
        this.deleteBOMItem = this.deleteBOMItem.bind(this);
    }

    deleteBOMItem(id){
        BOMItemService.deleteBOMItem(id).then( res => {
            this.setState({bOMItems: this.state.bOMItems.filter(bOMItem => bOMItem.bOMItemId !== id)});
        });
    }
    viewBOMItem(id){
        this.props.history.push(`/view-bOMItem/${id}`);
    }
    editBOMItem(id){
        this.props.history.push(`/add-bOMItem/${id}`);
    }

    componentDidMount(){
        BOMItemService.getBOMItems().then((res) => {
            this.setState({ bOMItems: res.data});
        });
    }

    addBOMItem(){
        this.props.history.push('/add-bOMItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BOMItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBOMItem}> Add BOMItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> Quantity </th>
                                    <th> ScrapPercent </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.bOMItems.map(
                                        bOMItem => 
                                        <tr key = {bOMItem.bOMItemId}>
                                             <td> { bOMItem.lineNumber } </td>
                                             <td> { bOMItem.quantity } </td>
                                             <td> { bOMItem.scrapPercent } </td>
                                             <td>
                                                 <button onClick={ () => this.editBOMItem(bOMItem.bOMItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBOMItem(bOMItem.bOMItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBOMItem(bOMItem.bOMItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBOMItemComponent
