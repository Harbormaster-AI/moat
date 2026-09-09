import React, { Component } from 'react'
import MatterService from '../services/MatterService'

class ListMatterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                matters: []
        }
        this.addMatter = this.addMatter.bind(this);
        this.editMatter = this.editMatter.bind(this);
        this.deleteMatter = this.deleteMatter.bind(this);
    }

    deleteMatter(id){
        MatterService.deleteMatter(id).then( res => {
            this.setState({matters: this.state.matters.filter(matter => matter.matterId !== id)});
        });
    }
    viewMatter(id){
        this.props.history.push(`/view-matter/${id}`);
    }
    editMatter(id){
        this.props.history.push(`/add-matter/${id}`);
    }

    componentDidMount(){
        MatterService.getMatters().then((res) => {
            this.setState({ matters: res.data});
        });
    }

    addMatter(){
        this.props.history.push('/add-matter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Matter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMatter}> Add Matter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> MatterName </th>
                                    <th> LeadCounsel </th>
                                    <th> MatterType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.matters.map(
                                        matter => 
                                        <tr key = {matter.matterId}>
                                             <td> { matter.matterName } </td>
                                             <td> { matter.leadCounsel } </td>
                                             <td> { matter.matterType } </td>
                                             <td> { matter.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMatter(matter.matterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMatter(matter.matterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMatter(matter.matterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMatterComponent
