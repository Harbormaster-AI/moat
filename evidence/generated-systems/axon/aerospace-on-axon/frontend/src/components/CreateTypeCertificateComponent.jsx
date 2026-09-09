import React, { Component } from 'react'
import TypeCertificateService from '../services/TypeCertificateService';

class CreateTypeCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                certificateNumber: '',
                authority: ''
        }
        this.changecertificateNumberHandler = this.changecertificateNumberHandler.bind(this);
        this.changeauthorityHandler = this.changeauthorityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TypeCertificateService.getTypeCertificateById(this.state.id).then( (res) =>{
                let typeCertificate = res.data;
                this.setState({
                    certificateNumber: typeCertificate.certificateNumber,
                    authority: typeCertificate.authority
                });
            });
        }        
    }
    saveOrUpdateTypeCertificate = (e) => {
        e.preventDefault();
        let typeCertificate = {
                typeCertificateId: this.state.id,
                certificateNumber: this.state.certificateNumber,
                authority: this.state.authority
            };
        console.log('typeCertificate => ' + JSON.stringify(typeCertificate));

        // step 5
        if(this.state.id === '_add'){
            typeCertificate.typeCertificateId=''
            TypeCertificateService.createTypeCertificate(typeCertificate).then(res =>{
                this.props.history.push('/typeCertificates');
            });
        }else{
            TypeCertificateService.updateTypeCertificate(typeCertificate).then( res => {
                this.props.history.push('/typeCertificates');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TypeCertificate</h3>
        }else{
            return <h3 className="text-center">Update TypeCertificate</h3>
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
                                            <label> certificateNumber:&emsp; </label>
                                                <input placeholder="certificateNumber" name="certificateNumber" className="form-control" value={this.state.certificateNumber} onChange={this.changecertificateNumberHandler}/>

                                            <label> authority:&emsp; </label>
                                                <input placeholder="authority" name="authority" className="form-control" value={this.state.authority} onChange={this.changeauthorityHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTypeCertificate}>Save</button>
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

export default CreateTypeCertificateComponent
