import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BackgroundCheckService } from '../../../services/BackgroundCheck.service';
import { BackgroundCheck } from '../../../models/BackgroundCheck';
import { SubBaseComponent } from '../../BackgroundCheck/sub.base.component';

@Component({
    selector: 'app-create-backgroundCheck',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBackgroundCheckComponent extends SubBaseComponent implements OnInit {

    title = 'Add BackgroundCheck';

    backgroundCheckForm: FormGroup;
    backgroundCheck: BackgroundCheck;

    constructor( http: HttpClient,
        private backgroundCheckService: BackgroundCheckService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status): void {
        this.backgroundCheckService
        .addBackgroundCheck(checkNumber, provider, completedDate, Candidate, Requisition, Report, Status)
            .subscribe(() => {
                this.router.navigate(['/indexBackgroundCheck']);
            });
    }

    ngOnInit(): void {
    }
}