import React, { Component } from 'react'
import BOMService from '../services/BOMService'

class ListBOMComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                bOMs: []
        }
        this.addBOM = this.addBOM.bind(this);
        this.editBOM = this.editBOM.bind(this);
        this.deleteBOM = this.deleteBOM.bind(this);
    }

    deleteBOM(id){
        BOMService.deleteBOM(id).then( res => {
            this.setState({bOMs: this.state.bOMs.filter(bOM => bOM.bOMId !== id)});
        });
    }
    viewBOM(id){
        this.props.history.push(`/view-bOM/${id}`);
    }
    editBOM(id){
        this.props.history.push(`/add-bOM/${id}`);
    }

    componentDidMount(){
        BOMService.getBOMs().then((res) => {
            this.setState({ bOMs: res.data});
        });
    }

    addBOM(){
        this.props.history.push('/add-bOM/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">BOM List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addBOM}> Add BOM</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BomNumber </th>
                                    <th> Revision </th>
                                    <th> EffectivityStart </th>
                                    <th> EffectivityEnd </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.bOMs.map(
                                        bOM => 
                                        <tr key = {bOM.bOMId}>
                                             <td> { bOM.bomNumber } </td>
                                             <td> { bOM.revision } </td>
                                             <td> { bOM.effectivityStart } </td>
                                             <td> { bOM.effectivityEnd } </td>
                                             <td> { bOM.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editBOM(bOM.bOMId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteBOM(bOM.bOMId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewBOM(bOM.bOMId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListBOMComponent
