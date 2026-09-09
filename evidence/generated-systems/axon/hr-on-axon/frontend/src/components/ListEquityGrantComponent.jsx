import React, { Component } from 'react'
import EquityGrantService from '../services/EquityGrantService'

class ListEquityGrantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                equityGrants: []
        }
        this.addEquityGrant = this.addEquityGrant.bind(this);
        this.editEquityGrant = this.editEquityGrant.bind(this);
        this.deleteEquityGrant = this.deleteEquityGrant.bind(this);
    }

    deleteEquityGrant(id){
        EquityGrantService.deleteEquityGrant(id).then( res => {
            this.setState({equityGrants: this.state.equityGrants.filter(equityGrant => equityGrant.equityGrantId !== id)});
        });
    }
    viewEquityGrant(id){
        this.props.history.push(`/view-equityGrant/${id}`);
    }
    editEquityGrant(id){
        this.props.history.push(`/add-equityGrant/${id}`);
    }

    componentDidMount(){
        EquityGrantService.getEquityGrants().then((res) => {
            this.setState({ equityGrants: res.data});
        });
    }

    addEquityGrant(){
        this.props.history.push('/add-equityGrant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EquityGrant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEquityGrant}> Add EquityGrant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> GrantId </th>
                                    <th> GrantedUnits </th>
                                    <th> VestingStart </th>
                                    <th> GrantType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.equityGrants.map(
                                        equityGrant => 
                                        <tr key = {equityGrant.equityGrantId}>
                                             <td> { equityGrant.grantId } </td>
                                             <td> { equityGrant.grantedUnits } </td>
                                             <td> { equityGrant.vestingStart } </td>
                                             <td> { equityGrant.grantType } </td>
                                             <td>
                                                 <button onClick={ () => this.editEquityGrant(equityGrant.equityGrantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEquityGrant(equityGrant.equityGrantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEquityGrant(equityGrant.equityGrantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEquityGrantComponent
