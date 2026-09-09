import React, { Component } from 'react'
import ProductionLineService from '../services/ProductionLineService'

class ListProductionLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productionLines: []
        }
        this.addProductionLine = this.addProductionLine.bind(this);
        this.editProductionLine = this.editProductionLine.bind(this);
        this.deleteProductionLine = this.deleteProductionLine.bind(this);
    }

    deleteProductionLine(id){
        ProductionLineService.deleteProductionLine(id).then( res => {
            this.setState({productionLines: this.state.productionLines.filter(productionLine => productionLine.productionLineId !== id)});
        });
    }
    viewProductionLine(id){
        this.props.history.push(`/view-productionLine/${id}`);
    }
    editProductionLine(id){
        this.props.history.push(`/add-productionLine/${id}`);
    }

    componentDidMount(){
        ProductionLineService.getProductionLines().then((res) => {
            this.setState({ productionLines: res.data});
        });
    }

    addProductionLine(){
        this.props.history.push('/add-productionLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductionLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductionLine}> Add ProductionLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LineCode </th>
                                    <th> LineType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productionLines.map(
                                        productionLine => 
                                        <tr key = {productionLine.productionLineId}>
                                             <td> { productionLine.name } </td>
                                             <td> { productionLine.lineCode } </td>
                                             <td> { productionLine.lineType } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductionLine(productionLine.productionLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductionLine(productionLine.productionLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductionLine(productionLine.productionLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductionLineComponent
