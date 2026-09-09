import React, { Component } from 'react'
import DimensionService from '../services/DimensionService'

class ListDimensionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dimensions: []
        }
        this.addDimension = this.addDimension.bind(this);
        this.editDimension = this.editDimension.bind(this);
        this.deleteDimension = this.deleteDimension.bind(this);
    }

    deleteDimension(id){
        DimensionService.deleteDimension(id).then( res => {
            this.setState({dimensions: this.state.dimensions.filter(dimension => dimension.dimensionId !== id)});
        });
    }
    viewDimension(id){
        this.props.history.push(`/view-dimension/${id}`);
    }
    editDimension(id){
        this.props.history.push(`/add-dimension/${id}`);
    }

    componentDidMount(){
        DimensionService.getDimensions().then((res) => {
            this.setState({ dimensions: res.data});
        });
    }

    addDimension(){
        this.props.history.push('/add-dimension/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Dimension List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDimension}> Add Dimension</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TypeTime </th>
                                    <th> DimensionType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dimensions.map(
                                        dimension => 
                                        <tr key = {dimension.dimensionId}>
                                             <td> { dimension.name } </td>
                                             <td> { dimension.typeTime } </td>
                                             <td> { dimension.dimensionType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDimension(dimension.dimensionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDimension(dimension.dimensionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDimension(dimension.dimensionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDimensionComponent
