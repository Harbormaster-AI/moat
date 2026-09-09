import React, { Component } from 'react'
import AdvertiserService from '../services/AdvertiserService';

class UpdateAdvertiserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                industry: '',
                website: ''
        }
        this.updateAdvertiser = this.updateAdvertiser.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeindustryHandler = this.changeindustryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    componentDidMount(){
        AdvertiserService.getAdvertiserById(this.state.id).then( (res) =>{
            let advertiser = res.data;
            this.setState({
                name: advertiser.name,
                legalName: advertiser.legalName,
                industry: advertiser.industry,
                website: advertiser.website
            });
        });
    }

    updateAdvertiser = (e) => {
        e.preventDefault();
        let advertiser = {
            advertiserId: this.state.id,
            name: this.state.name,
            legalName: this.state.legalName,
            industry: this.state.industry,
            website: this.state.website
        };
        console.log('advertiser => ' + JSON.stringify(advertiser));
        console.log('id => ' + JSON.stringify(this.state.id));
        AdvertiserService.updateAdvertiser(advertiser).then( res => {
            this.props.history.push('/advertisers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changeindustryHandler= (event) => {
        this.setState({industry: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/advertisers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Advertiser</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> legalName: </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> industry: </label>
                                                <input placeholder="industry" name="industry" className="form-control" value={this.state.industry} onChange={this.changeindustryHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAdvertiser}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateAdvertiserComponent
