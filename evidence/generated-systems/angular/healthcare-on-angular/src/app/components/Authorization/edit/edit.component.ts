import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuthorizationService } from '../../../services/Authorization.service';
import { SubBaseComponent } from '../../Authorization/sub.base.component';


@Component({
    selector: 'app-edit-authorization',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAuthorizationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Authorization';

    authorizationForm: FormGroup;
    authorization: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AuthorizationService,
        private fb: FormBuilder
) {
        super(http);
        this.authorizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  authNumber: ['', Validators.required],
      requestedService: ['', Validators.required],
      Coverage: ['', ],
      Order: ['', ],
      Status: ['', ]
        });
    }

    
    updateAuthorization(authNumber, requestedService, Coverage, Order, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAuthorization(authNumber, requestedService, Coverage, Order, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAuthorization']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAuthorization(params['id']).subscribe(res => {
                this.authorization = res;
            });
        });
    }
}