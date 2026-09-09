import React, { Component } from 'react'
import TypeCertificateService from '../services/TypeCertificateService';

class UpdateTypeCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                certificateNumber: '',
                authority: ''
        }
        this.updateTypeCertificate = this.updateTypeCertificate.bind(this);

        this.changecertificateNumberHandler = this.changecertificateNumberHandler.bind(this);
        this.changeauthorityHandler = this.changeauthorityHandler.bind(this);
    }

    componentDidMount(){
        TypeCertificateService.getTypeCertificateById(this.state.id).then( (res) =>{
            let typeCertificate = res.data;
            this.setState({
                certificateNumber: typeCertificate.certificateNumber,
                authority: typeCertificate.authority
            });
        });
    }

    updateTypeCertificate = (e) => {
        e.preventDefault();
        let typeCertificate = {
            typeCertificateId: this.state.id,
            certificateNumber: this.state.certificateNumber,
            authority: this.state.authority
        };
        console.log('typeCertificate => ' + JSON.stringify(typeCertificate));
        console.log('id => ' + JSON.stringify(this.state.id));
        TypeCertificateService.updateTypeCertificate(typeCertificate).then( res => {
            this.props.history.push('/typeCertificates');
        });
    }

    changecertificateNumberHandler= (event) => {
        this.setState({certificateNumber: event.target.value});
    }
    changeauthorityHandler= (event) => {
        this.setState({authority: event.target.value});
    }

    cancel(){
        this.props.history.push('/typeCertificates');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TypeCertificate</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> certificateNumber: </label>
                                                <input placeholder="certificateNumber" name="certificateNumber" className="form-control" value={this.state.certificateNumber} onChange={this.changecertificateNumberHandler}/>

                                            <label> authority: </label>
                                                <input placeholder="authority" name="authority" className="form-control" value={this.state.authority} onChange={this.changeauthorityHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTypeCertificate}>Save</button>
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

export default UpdateTypeCertificateComponent
