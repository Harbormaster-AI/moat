import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';
import { SoftwareLoad } from '../../../models/SoftwareLoad';
import { SubBaseComponent } from '../../SoftwareLoad/sub.base.component';

@Component({
    selector: 'app-create-softwareLoad',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSoftwareLoadComponent extends SubBaseComponent implements OnInit {

    title = 'Add SoftwareLoad';

    softwareLoadForm: FormGroup;
    softwareLoad: SoftwareLoad;

    constructor( http: HttpClient,
        private softwareLoadService: SoftwareLoadService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.softwareLoadForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  version: ['', Validators.required],
      ConnectedAircraft: ['', ],
      AvionicsSuite: ['', ],
      LoadType: ['', ]
        });
    }

    
    addSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType): void {
        this.softwareLoadService
        .addSoftwareLoad(version, ConnectedAircraft, AvionicsSuite, LoadType)
            .subscribe(() => {
                this.router.navigate(['/indexSoftwareLoad']);
            });
    }

    ngOnInit(): void {
    }
}