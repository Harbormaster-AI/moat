import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ScreeningService } from '../../../services/Screening.service';
import { Screening } from '../../../models/Screening';
import { SubBaseComponent } from '../../Screening/sub.base.component';

@Component({
    selector: 'app-create-screening',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateScreeningComponent extends SubBaseComponent implements OnInit {

    title = 'Add Screening';

    screeningForm: FormGroup;
    screening: Screening;

    constructor( http: HttpClient,
        private screeningService: ScreeningService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.screeningForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      completedDate: ['', Validators.required],
      Application: ['', ],
      Status: ['', ]
        });
    }

    
    addScreening(name, completedDate, Application, Status): void {
        this.screeningService
        .addScreening(name, completedDate, Application, Status)
            .subscribe(() => {
                this.router.navigate(['/indexScreening']);
            });
    }

    ngOnInit(): void {
    }
}