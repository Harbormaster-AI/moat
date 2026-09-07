
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateImagingReportComponent } from './create.component';
import { ImagingReportService } from '../../../services/ImagingReport.service';
import { Router } from '@angular/router';

describe('CreateImagingReportComponent', () => {
  let component: CreateImagingReportComponent;
  let fixture: ComponentFixture<CreateImagingReportComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateImagingReportComponent
      ],
      providers: [
        ImagingReportService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateImagingReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});