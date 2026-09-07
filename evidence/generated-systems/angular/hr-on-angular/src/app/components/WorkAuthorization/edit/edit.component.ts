import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';
import { SubBaseComponent } from '../../WorkAuthorization/sub.base.component';


@Component({
    selector: 'app-edit-workAuthorization',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWorkAuthorizationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit WorkAuthorization';

    workAuthorizationForm: FormGroup;
    workAuthorization: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WorkAuthorizationService,
        private fb: FormBuilder
) {
        super(http);
        this.workAuthorizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  country: ['', Validators.required],
      expirationDate: ['', Validators.required],
      Employee: ['', ],
      Documents: ['', ],
      Status: ['', ]
        });
    }

    
    updateWorkAuthorization(country, expirationDate, Employee, Documents, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWorkAuthorization(country, expirationDate, Employee, Documents, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWorkAuthorization']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWorkAuthorization(params['id']).subscribe(res => {
                this.workAuthorization = res;
            });
        });
    }
}