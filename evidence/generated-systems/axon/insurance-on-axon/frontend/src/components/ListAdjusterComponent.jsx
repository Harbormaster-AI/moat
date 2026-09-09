import React, { Component } from 'react'
import AdjusterService from '../services/AdjusterService'

class ListAdjusterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                adjusters: []
        }
        this.addAdjuster = this.addAdjuster.bind(this);
        this.editAdjuster = this.editAdjuster.bind(this);
        this.deleteAdjuster = this.deleteAdjuster.bind(this);
    }

    deleteAdjuster(id){
        AdjusterService.deleteAdjuster(id).then( res => {
            this.setState({adjusters: this.state.adjusters.filter(adjuster => adjuster.adjusterId !== id)});
        });
    }
    viewAdjuster(id){
        this.props.history.push(`/view-adjuster/${id}`);
    }
    editAdjuster(id){
        this.props.history.push(`/add-adjuster/${id}`);
    }

    componentDidMount(){
        AdjusterService.getAdjusters().then((res) => {
            this.setState({ adjusters: res.data});
        });
    }

    addAdjuster(){
        this.props.history.push('/add-adjuster/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Adjuster List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAdjuster}> Add Adjuster</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> LicenseNumber </th>
                                    <th> AdjusterType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.adjusters.map(
                                        adjuster => 
                                        <tr key = {adjuster.adjusterId}>
                                             <td> { adjuster.firstName } </td>
                                             <td> { adjuster.lastName } </td>
                                             <td> { adjuster.licenseNumber } </td>
                                             <td> { adjuster.adjusterType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAdjuster(adjuster.adjusterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAdjuster(adjuster.adjusterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAdjuster(adjuster.adjusterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAdjusterComponent
