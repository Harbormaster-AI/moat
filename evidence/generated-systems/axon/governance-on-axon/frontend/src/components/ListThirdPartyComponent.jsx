import React, { Component } from 'react'
import ThirdPartyService from '../services/ThirdPartyService'

class ListThirdPartyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                thirdPartys: []
        }
        this.addThirdParty = this.addThirdParty.bind(this);
        this.editThirdParty = this.editThirdParty.bind(this);
        this.deleteThirdParty = this.deleteThirdParty.bind(this);
    }

    deleteThirdParty(id){
        ThirdPartyService.deleteThirdParty(id).then( res => {
            this.setState({thirdPartys: this.state.thirdPartys.filter(thirdParty => thirdParty.thirdPartyId !== id)});
        });
    }
    viewThirdParty(id){
        this.props.history.push(`/view-thirdParty/${id}`);
    }
    editThirdParty(id){
        this.props.history.push(`/add-thirdParty/${id}`);
    }

    componentDidMount(){
        ThirdPartyService.getThirdPartys().then((res) => {
            this.setState({ thirdPartys: res.data});
        });
    }

    addThirdParty(){
        this.props.history.push('/add-thirdParty/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ThirdParty List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addThirdParty}> Add ThirdParty</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Country </th>
                                    <th> ContactEmail </th>
                                    <th> ThirdPartyType </th>
                                    <th> Criticality </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.thirdPartys.map(
                                        thirdParty => 
                                        <tr key = {thirdParty.thirdPartyId}>
                                             <td> { thirdParty.name } </td>
                                             <td> { thirdParty.country } </td>
                                             <td> { thirdParty.contactEmail } </td>
                                             <td> { thirdParty.thirdPartyType } </td>
                                             <td> { thirdParty.criticality } </td>
                                             <td>
                                                 <button onClick={ () => this.editThirdParty(thirdParty.thirdPartyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteThirdParty(thirdParty.thirdPartyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewThirdParty(thirdParty.thirdPartyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListThirdPartyComponent
