import React, { Component } from 'react'
import UoMConversionService from '../services/UoMConversionService'

class ListUoMConversionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                uoMConversions: []
        }
        this.addUoMConversion = this.addUoMConversion.bind(this);
        this.editUoMConversion = this.editUoMConversion.bind(this);
        this.deleteUoMConversion = this.deleteUoMConversion.bind(this);
    }

    deleteUoMConversion(id){
        UoMConversionService.deleteUoMConversion(id).then( res => {
            this.setState({uoMConversions: this.state.uoMConversions.filter(uoMConversion => uoMConversion.uoMConversionId !== id)});
        });
    }
    viewUoMConversion(id){
        this.props.history.push(`/view-uoMConversion/${id}`);
    }
    editUoMConversion(id){
        this.props.history.push(`/add-uoMConversion/${id}`);
    }

    componentDidMount(){
        UoMConversionService.getUoMConversions().then((res) => {
            this.setState({ uoMConversions: res.data});
        });
    }

    addUoMConversion(){
        this.props.history.push('/add-uoMConversion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">UoMConversion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addUoMConversion}> Add UoMConversion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Factor </th>
                                    <th> Precision </th>
                                    <th> FromUnit </th>
                                    <th> ToUnit </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.uoMConversions.map(
                                        uoMConversion => 
                                        <tr key = {uoMConversion.uoMConversionId}>
                                             <td> { uoMConversion.factor } </td>
                                             <td> { uoMConversion.precision } </td>
                                             <td> { uoMConversion.fromUnit } </td>
                                             <td> { uoMConversion.toUnit } </td>
                                             <td>
                                                 <button onClick={ () => this.editUoMConversion(uoMConversion.uoMConversionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteUoMConversion(uoMConversion.uoMConversionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewUoMConversion(uoMConversion.uoMConversionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListUoMConversionComponent
