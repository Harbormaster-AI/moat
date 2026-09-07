import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ApplicationService } from '../../../services/Application.service';
import { SubBaseComponent } from '../../Application/sub.base.component';


@Component({
    selector: 'app-edit-application',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Application';

    applicationForm: FormGroup;
    application: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ApplicationService,
        private fb: FormBuilder
) {
        super(http);
        this.applicationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  applicationNumber: ['', Validators.required],
      submissionDate: ['', Validators.required],
      Customer: ['', ],
      Product: ['', ],
      Distributor: ['', ],
      Quotes: ['', ],
      SelectedQuote: ['', ],
      Status: ['', ]
        });
    }

    
    updateApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexApplication']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getApplication(params['id']).subscribe(res => {
                this.application = res;
            });
        });
    }
}