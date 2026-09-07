import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BackgroundCheckService } from '../../../services/BackgroundCheck.service';
import { SubBaseComponent } from '../../BackgroundCheck/sub.base.component';


@Component({
    selector: 'app-edit-backgroundCheck',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBackgroundCheckComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BackgroundCheck';

    backgroundCheckForm: FormGroup;
    backgroundCheck: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BackgroundCheckService,
        private fb: FormBuilder
) {
        super(http);
        this.backgroundCheckForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  checkNumber: ['', Validators.required],
      provider: ['', Validators.required],
      completedDate: ['', Validators.required],
      Candidate: ['', ],
      Requisition: ['', ],
      Report: ['', ],
      Status: ['', ]
        });
    }

    
    updateBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBackgroundCheck']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBackgroundCheck(params['id']).subscribe(res => {
                this.backgroundCheck = res;
            });
        });
    }
}