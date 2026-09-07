import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataProviderService } from '../../../services/DataProvider.service';
import { DataProvider } from '../../../models/DataProvider';
import { SubBaseComponent } from '../../DataProvider/sub.base.component';

@Component({
    selector: 'app-create-dataProvider',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataProviderComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataProvider';

    dataProviderForm: FormGroup;
    dataProvider: DataProvider;

    constructor( http: HttpClient,
        private dataProviderService: DataProviderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataProviderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      AudienceSegments: ['', ],
      ProviderType: ['', ]
        });
    }

    
    addDataProvider(name, website, AudienceSegments, ProviderType): void {
        this.dataProviderService
        .addDataProvider(name, website, AudienceSegments, ProviderType)
            .subscribe(() => {
                this.router.navigate(['/indexDataProvider']);
            });
    }

    ngOnInit(): void {
    }
}