import React, { Component } from 'react'
import AuditWorkpaperService from '../services/AuditWorkpaperService';

class CreateAuditWorkpaperComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                workpaperRef: '',
                subject: '',
                workpaperUrl: ''
        }
        this.changeworkpaperRefHandler = this.changeworkpaperRefHandler.bind(this);
        this.changesubjectHandler = this.changesubjectHandler.bind(this);
        this.changeworkpaperUrlHandler = this.changeworkpaperUrlHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AuditWorkpaperService.getAuditWorkpaperById(this.state.id).then( (res) =>{
                let auditWorkpaper = res.data;
                this.setState({
                    workpaperRef: auditWorkpaper.workpaperRef,
                    subject: auditWorkpaper.subject,
                    workpaperUrl: auditWorkpaper.workpaperUrl
                });
            });
        }        
    }
    saveOrUpdateAuditWorkpaper = (e) => {
        e.preventDefault();
        let auditWorkpaper = {
                auditWorkpaperId: this.state.id,
                workpaperRef: this.state.workpaperRef,
                subject: this.state.subject,
                workpaperUrl: this.state.workpaperUrl
            };
        console.log('auditWorkpaper => ' + JSON.stringify(auditWorkpaper));

        // step 5
        if(this.state.id === '_add'){
            auditWorkpaper.auditWorkpaperId=''
            AuditWorkpaperService.createAuditWorkpaper(auditWorkpaper).then(res =>{
                this.props.history.push('/auditWorkpapers');
            });
        }else{
            AuditWorkpaperService.updateAuditWorkpaper(auditWorkpaper).then( res => {
                this.props.history.push('/auditWorkpapers');
            });
        }
    }
    
    changeworkpaperRefHandler= (event) => {
        this.setState({workpaperRef: event.target.value});
    }
    changesubjectHandler= (event) => {
        this.setState({subject: event.target.value});
    }
    changeworkpaperUrlHandler= (event) => {
        this.setState({workpaperUrl: event.target.value});
    }

    cancel(){
        this.props.history.push('/auditWorkpapers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AuditWorkpaper</h3>
        }else{
            return <h3 className="text-center">Update AuditWorkpaper</h3>
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
                                            <label> workpaperRef:&emsp; </label>
                                                <input placeholder="workpaperRef" name="workpaperRef" className="form-control" value={this.state.workpaperRef} onChange={this.changeworkpaperRefHandler}/>

                                            <label> subject:&emsp; </label>
                                                <input placeholder="subject" name="subject" className="form-control" value={this.state.subject} onChange={this.changesubjectHandler}/>

                                            <label> workpaperUrl:&emsp; </label>
                                                <input placeholder="workpaperUrl" name="workpaperUrl" className="form-control" value={this.state.workpaperUrl} onChange={this.changeworkpaperUrlHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAuditWorkpaper}>Save</button>
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

export default CreateAuditWorkpaperComponent
