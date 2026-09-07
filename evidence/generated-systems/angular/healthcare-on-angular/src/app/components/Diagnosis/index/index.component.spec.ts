
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDiagnosisComponent } from './index.component';
import { DiagnosisService } from '../../../services/Diagnosis.service';

describe('IndexDiagnosisComponent', () => {
  let component: IndexDiagnosisComponent;
  let fixture: ComponentFixture<IndexDiagnosisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDiagnosisComponent
      ],
      providers: [
        DiagnosisService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDiagnosisComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});