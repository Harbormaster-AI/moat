import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreativeFileService } from '../../../services/CreativeFile.service';
import { CreativeFile } from '../../../models/CreativeFile';
import { SubBaseComponent } from '../../CreativeFile/sub.base.component';

@Component({
    selector: 'app-create-creativeFile',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCreativeFileComponent extends SubBaseComponent implements OnInit {

    title = 'Add CreativeFile';

    creativeFileForm: FormGroup;
    creativeFile: CreativeFile;

    constructor( http: HttpClient,
        private creativeFileService: CreativeFileService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.creativeFileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  uri: ['', Validators.required],
      fileSizeKB: ['', Validators.required],
      mimeType: ['', Validators.required],
      checksum: ['', Validators.required],
      CreativeAsset: ['', ]
        });
    }

    
    addCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset): void {
        this.creativeFileService
        .addCreativeFile(uri, fileSizeKB, mimeType, checksum, CreativeAsset)
            .subscribe(() => {
                this.router.navigate(['/indexCreativeFile']);
            });
    }

    ngOnInit(): void {
    }
}