import React, { Component } from 'react'
import AdvertiserService from '../services/AdvertiserService'

class ListAdvertiserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                advertisers: []
        }
        this.addAdvertiser = this.addAdvertiser.bind(this);
        this.editAdvertiser = this.editAdvertiser.bind(this);
        this.deleteAdvertiser = this.deleteAdvertiser.bind(this);
    }

    deleteAdvertiser(id){
        AdvertiserService.deleteAdvertiser(id).then( res => {
            this.setState({advertisers: this.state.advertisers.filter(advertiser => advertiser.advertiserId !== id)});
        });
    }
    viewAdvertiser(id){
        this.props.history.push(`/view-advertiser/${id}`);
    }
    editAdvertiser(id){
        this.props.history.push(`/add-advertiser/${id}`);
    }

    componentDidMount(){
        AdvertiserService.getAdvertisers().then((res) => {
            this.setState({ advertisers: res.data});
        });
    }

    addAdvertiser(){
        this.props.history.push('/add-advertiser/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Advertiser List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAdvertiser}> Add Advertiser</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> Industry </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.advertisers.map(
                                        advertiser => 
                                        <tr key = {advertiser.advertiserId}>
                                             <td> { advertiser.name } </td>
                                             <td> { advertiser.legalName } </td>
                                             <td> { advertiser.industry } </td>
                                             <td> { advertiser.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editAdvertiser(advertiser.advertiserId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAdvertiser(advertiser.advertiserId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAdvertiser(advertiser.advertiserId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAdvertiserComponent
