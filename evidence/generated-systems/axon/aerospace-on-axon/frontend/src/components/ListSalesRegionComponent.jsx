import React, { Component } from 'react'
import SalesRegionService from '../services/SalesRegionService'

class ListSalesRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                salesRegions: []
        }
        this.addSalesRegion = this.addSalesRegion.bind(this);
        this.editSalesRegion = this.editSalesRegion.bind(this);
        this.deleteSalesRegion = this.deleteSalesRegion.bind(this);
    }

    deleteSalesRegion(id){
        SalesRegionService.deleteSalesRegion(id).then( res => {
            this.setState({salesRegions: this.state.salesRegions.filter(salesRegion => salesRegion.salesRegionId !== id)});
        });
    }
    viewSalesRegion(id){
        this.props.history.push(`/view-salesRegion/${id}`);
    }
    editSalesRegion(id){
        this.props.history.push(`/add-salesRegion/${id}`);
    }

    componentDidMount(){
        SalesRegionService.getSalesRegions().then((res) => {
            this.setState({ salesRegions: res.data});
        });
    }

    addSalesRegion(){
        this.props.history.push('/add-salesRegion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SalesRegion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSalesRegion}> Add SalesRegion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> RegionCode </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.salesRegions.map(
                                        salesRegion => 
                                        <tr key = {salesRegion.salesRegionId}>
                                             <td> { salesRegion.name } </td>
                                             <td> { salesRegion.regionCode } </td>
                                             <td>
                                                 <button onClick={ () => this.editSalesRegion(salesRegion.salesRegionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSalesRegion(salesRegion.salesRegionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSalesRegion(salesRegion.salesRegionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSalesRegionComponent
