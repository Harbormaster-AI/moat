import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';
import { SoftwareUpdate } from '../../../models/SoftwareUpdate';
import { SubBaseComponent } from '../../SoftwareUpdate/sub.base.component';

@Component({
    selector: 'app-create-softwareUpdate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSoftwareUpdateComponent extends SubBaseComponent implements OnInit {

    title = 'Add SoftwareUpdate';

    softwareUpdateForm: FormGroup;
    softwareUpdate: SoftwareUpdate;

    constructor( http: HttpClient,
        private softwareUpdateService: SoftwareUpdateService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.softwareUpdateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  version: ['', Validators.required],
      appliedDate: ['', Validators.required],
      Device: ['', ],
      UpdateType: ['', ]
        });
    }

    
    addSoftwareUpdate(version, appliedDate, Device, UpdateType): void {
        this.softwareUpdateService
        .addSoftwareUpdate(version, appliedDate, Device, UpdateType)
            .subscribe(() => {
                this.router.navigate(['/indexSoftwareUpdate']);
            });
    }

    ngOnInit(): void {
    }
}