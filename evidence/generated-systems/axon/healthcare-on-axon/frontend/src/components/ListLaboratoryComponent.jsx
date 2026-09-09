import React, { Component } from 'react'
import LaboratoryService from '../services/LaboratoryService'

class ListLaboratoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                laboratorys: []
        }
        this.addLaboratory = this.addLaboratory.bind(this);
        this.editLaboratory = this.editLaboratory.bind(this);
        this.deleteLaboratory = this.deleteLaboratory.bind(this);
    }

    deleteLaboratory(id){
        LaboratoryService.deleteLaboratory(id).then( res => {
            this.setState({laboratorys: this.state.laboratorys.filter(laboratory => laboratory.laboratoryId !== id)});
        });
    }
    viewLaboratory(id){
        this.props.history.push(`/view-laboratory/${id}`);
    }
    editLaboratory(id){
        this.props.history.push(`/add-laboratory/${id}`);
    }

    componentDidMount(){
        LaboratoryService.getLaboratorys().then((res) => {
            this.setState({ laboratorys: res.data});
        });
    }

    addLaboratory(){
        this.props.history.push('/add-laboratory/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Laboratory List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLaboratory}> Add Laboratory</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> CliaNumber </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.laboratorys.map(
                                        laboratory => 
                                        <tr key = {laboratory.laboratoryId}>
                                             <td> { laboratory.name } </td>
                                             <td> { laboratory.cliaNumber } </td>
                                             <td>
                                                 <button onClick={ () => this.editLaboratory(laboratory.laboratoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLaboratory(laboratory.laboratoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLaboratory(laboratory.laboratoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLaboratoryComponent
