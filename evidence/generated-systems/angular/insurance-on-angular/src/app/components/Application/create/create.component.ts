import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApplicationService } from '../../../services/Application.service';
import { Application } from '../../../models/Application';
import { SubBaseComponent } from '../../Application/sub.base.component';

@Component({
    selector: 'app-create-application',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Application';

    applicationForm: FormGroup;
    application: Application;

    constructor( http: HttpClient,
        private applicationService: ApplicationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status): void {
        this.applicationService
        .addApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status)
            .subscribe(() => {
                this.router.navigate(['/indexApplication']);
            });
    }

    ngOnInit(): void {
    }
}