import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ImagingReportService } from '../../../services/ImagingReport.service';
import { ImagingReport } from '../../../models/ImagingReport';
import { SubBaseComponent } from '../../ImagingReport/sub.base.component';

@Component({
    selector: 'app-create-imagingReport',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateImagingReportComponent extends SubBaseComponent implements OnInit {

    title = 'Add ImagingReport';

    imagingReportForm: FormGroup;
    imagingReport: ImagingReport;

    constructor( http: HttpClient,
        private imagingReportService: ImagingReportService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.imagingReportForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reportNumber: ['', Validators.required],
      impression: ['', Validators.required],
      reportedDate: ['', Validators.required],
      ImagingOrder: ['', ],
      Clinician: ['', ],
      Encounter: ['', ],
      ImagingCenter: ['', ],
      Status: ['', ]
        });
    }

    
    addImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status): void {
        this.imagingReportService
        .addImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status)
            .subscribe(() => {
                this.router.navigate(['/indexImagingReport']);
            });
    }

    ngOnInit(): void {
    }
}