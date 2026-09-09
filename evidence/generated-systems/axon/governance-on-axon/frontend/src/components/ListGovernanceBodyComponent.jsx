import React, { Component } from 'react'
import GovernanceBodyService from '../services/GovernanceBodyService'

class ListGovernanceBodyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                governanceBodys: []
        }
        this.addGovernanceBody = this.addGovernanceBody.bind(this);
        this.editGovernanceBody = this.editGovernanceBody.bind(this);
        this.deleteGovernanceBody = this.deleteGovernanceBody.bind(this);
    }

    deleteGovernanceBody(id){
        GovernanceBodyService.deleteGovernanceBody(id).then( res => {
            this.setState({governanceBodys: this.state.governanceBodys.filter(governanceBody => governanceBody.governanceBodyId !== id)});
        });
    }
    viewGovernanceBody(id){
        this.props.history.push(`/view-governanceBody/${id}`);
    }
    editGovernanceBody(id){
        this.props.history.push(`/add-governanceBody/${id}`);
    }

    componentDidMount(){
        GovernanceBodyService.getGovernanceBodys().then((res) => {
            this.setState({ governanceBodys: res.data});
        });
    }

    addGovernanceBody(){
        this.props.history.push('/add-governanceBody/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GovernanceBody List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGovernanceBody}> Add GovernanceBody</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> CharterUrl </th>
                                    <th> Chair </th>
                                    <th> BodyType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.governanceBodys.map(
                                        governanceBody => 
                                        <tr key = {governanceBody.governanceBodyId}>
                                             <td> { governanceBody.name } </td>
                                             <td> { governanceBody.charterUrl } </td>
                                             <td> { governanceBody.chair } </td>
                                             <td> { governanceBody.bodyType } </td>
                                             <td>
                                                 <button onClick={ () => this.editGovernanceBody(governanceBody.governanceBodyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGovernanceBody(governanceBody.governanceBodyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGovernanceBody(governanceBody.governanceBodyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListGovernanceBodyComponent
