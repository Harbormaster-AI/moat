import React, { Component } from 'react'
import CertificationService from '../services/CertificationService';

class UpdateCertificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                issuer: '',
                validFrom: '',
                validTo: '',
                credentialId: ''
        }
        this.updateCertification = this.updateCertification.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeissuerHandler = this.changeissuerHandler.bind(this);
        this.changevalidFromHandler = this.changevalidFromHandler.bind(this);
        this.changevalidToHandler = this.changevalidToHandler.bind(this);
        this.changecredentialIdHandler = this.changecredentialIdHandler.bind(this);
    }

    componentDidMount(){
        CertificationService.getCertificationById(this.state.id).then( (res) =>{
            let certification = res.data;
            this.setState({
                name: certification.name,
                issuer: certification.issuer,
                validFrom: certification.validFrom,
                validTo: certification.validTo,
                credentialId: certification.credentialId
            });
        });
    }

    updateCertification = (e) => {
        e.preventDefault();
        let certification = {
            certificationId: this.state.id,
            name: this.state.name,
            issuer: this.state.issuer,
            validFrom: this.state.validFrom,
            validTo: this.state.validTo,
            credentialId: this.state.credentialId
        };
        console.log('certification => ' + JSON.stringify(certification));
        console.log('id => ' + JSON.stringify(this.state.id));
        CertificationService.updateCertification(certification).then( res => {
            this.props.history.push('/certifications');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeissuerHandler= (event) => {
        this.setState({issuer: event.target.value});
    }
    changevalidFromHandler= (event) => {
        this.setState({validFrom: event.target.value});
    }
    changevalidToHandler= (event) => {
        this.setState({validTo: event.target.value});
    }
    changecredentialIdHandler= (event) => {
        this.setState({credentialId: event.target.value});
    }

    cancel(){
        this.props.history.push('/certifications');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Certification</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> issuer: </label>
                                                <input placeholder="issuer" name="issuer" className="form-control" value={this.state.issuer} onChange={this.changeissuerHandler}/>

                                            <label> validFrom: </label>
                                                <input type="date" placeholder="validFrom" name="validFrom" className="form-control" value={this.state.validFrom} onChange={this.changevalidFromHandler}/>

                                            <label> validTo: </label>
                                                <input type="date" placeholder="validTo" name="validTo" className="form-control" value={this.state.validTo} onChange={this.changevalidToHandler}/>

                                            <label> credentialId: </label>
                                                <input placeholder="credentialId" name="credentialId" className="form-control" value={this.state.credentialId} onChange={this.changecredentialIdHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCertification}>Save</button>
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

export default UpdateCertificationComponent
