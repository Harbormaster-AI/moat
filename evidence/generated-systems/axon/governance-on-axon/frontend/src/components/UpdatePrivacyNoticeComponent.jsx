import React, { Component } from 'react'
import PrivacyNoticeService from '../services/PrivacyNoticeService';

class UpdatePrivacyNoticeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                audience: '',
                versionLabel: '',
                publicationDate: '',
                publicationUrl: '',
                status: ''
        }
        this.updatePrivacyNotice = this.updatePrivacyNotice.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeaudienceHandler = this.changeaudienceHandler.bind(this);
        this.changeversionLabelHandler = this.changeversionLabelHandler.bind(this);
        this.changepublicationDateHandler = this.changepublicationDateHandler.bind(this);
        this.changepublicationUrlHandler = this.changepublicationUrlHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PrivacyNoticeService.getPrivacyNoticeById(this.state.id).then( (res) =>{
            let privacyNotice = res.data;
            this.setState({
                title: privacyNotice.title,
                audience: privacyNotice.audience,
                versionLabel: privacyNotice.versionLabel,
                publicationDate: privacyNotice.publicationDate,
                publicationUrl: privacyNotice.publicationUrl,
                status: privacyNotice.status
            });
        });
    }

    updatePrivacyNotice = (e) => {
        e.preventDefault();
        let privacyNotice = {
            privacyNoticeId: this.state.id,
            title: this.state.title,
            audience: this.state.audience,
            versionLabel: this.state.versionLabel,
            publicationDate: this.state.publicationDate,
            publicationUrl: this.state.publicationUrl,
            status: this.state.status
        };
        console.log('privacyNotice => ' + JSON.stringify(privacyNotice));
        console.log('id => ' + JSON.stringify(this.state.id));
        PrivacyNoticeService.updatePrivacyNotice(privacyNotice).then( res => {
            this.props.history.push('/privacyNotices');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeaudienceHandler= (event) => {
        this.setState({audience: event.target.value});
    }
    changeversionLabelHandler= (event) => {
        this.setState({versionLabel: event.target.value});
    }
    changepublicationDateHandler= (event) => {
        this.setState({publicationDate: event.target.value});
    }
    changepublicationUrlHandler= (event) => {
        this.setState({publicationUrl: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/privacyNotices');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PrivacyNotice</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> audience: </label>
                                                <input placeholder="audience" name="audience" className="form-control" value={this.state.audience} onChange={this.changeaudienceHandler}/>

                                            <label> versionLabel: </label>
                                                <input placeholder="versionLabel" name="versionLabel" className="form-control" value={this.state.versionLabel} onChange={this.changeversionLabelHandler}/>

                                            <label> publicationDate: </label>
                                                <input type="date" placeholder="publicationDate" name="publicationDate" className="form-control" value={this.state.publicationDate} onChange={this.changepublicationDateHandler}/>

                                            <label> publicationUrl: </label>
                                                <input placeholder="publicationUrl" name="publicationUrl" className="form-control" value={this.state.publicationUrl} onChange={this.changepublicationUrlHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          InReview
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePrivacyNotice}>Save</button>
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

export default UpdatePrivacyNoticeComponent
