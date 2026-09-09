import React, { Component } from 'react'
import DSPService from '../services/DSPService'

class ListDSPComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dSPs: []
        }
        this.addDSP = this.addDSP.bind(this);
        this.editDSP = this.editDSP.bind(this);
        this.deleteDSP = this.deleteDSP.bind(this);
    }

    deleteDSP(id){
        DSPService.deleteDSP(id).then( res => {
            this.setState({dSPs: this.state.dSPs.filter(dSP => dSP.dSPId !== id)});
        });
    }
    viewDSP(id){
        this.props.history.push(`/view-dSP/${id}`);
    }
    editDSP(id){
        this.props.history.push(`/add-dSP/${id}`);
    }

    componentDidMount(){
        DSPService.getDSPs().then((res) => {
            this.setState({ dSPs: res.data});
        });
    }

    addDSP(){
        this.props.history.push('/add-dSP/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DSP List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDSP}> Add DSP</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Website </th>
                                    <th> Region </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dSPs.map(
                                        dSP => 
                                        <tr key = {dSP.dSPId}>
                                             <td> { dSP.name } </td>
                                             <td> { dSP.website } </td>
                                             <td> { dSP.region } </td>
                                             <td>
                                                 <button onClick={ () => this.editDSP(dSP.dSPId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDSP(dSP.dSPId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDSP(dSP.dSPId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDSPComponent
