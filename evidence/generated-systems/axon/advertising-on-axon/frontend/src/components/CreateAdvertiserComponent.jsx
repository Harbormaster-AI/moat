import React, { Component } from 'react'
import AdvertiserService from '../services/AdvertiserService';

class CreateAdvertiserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                industry: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeindustryHandler = this.changeindustryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateAdvertiser = (e) => {
        e.preventDefault();
        let advertiser = {
                advertiserId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                industry: this.state.industry,
                website: this.state.website
            };
        console.log('advertiser => ' + JSON.stringify(advertiser));

        // step 5
        if(this.state.id === '_add'){
            advertiser.advertiserId=''
            AdvertiserService.createAdvertiser(advertiser).then(res =>{
                this.props.history.push('/advertisers');
            });
        }else{
            AdvertiserService.updateAdvertiser(advertiser).then( res => {
                this.props.history.push('/advertisers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Advertiser</h3>
        }else{
            return <h3 className="text-center">Update Advertiser</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> legalName:&emsp; </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> industry:&emsp; </label>
                                                <input placeholder="industry" name="industry" className="form-control" value={this.state.industry} onChange={this.changeindustryHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAdvertiser}>Save</button>
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

export default CreateAdvertiserComponent
