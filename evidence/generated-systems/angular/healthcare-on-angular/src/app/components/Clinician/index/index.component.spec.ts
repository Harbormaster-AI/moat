
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexClinicianComponent } from './index.component';
import { ClinicianService } from '../../../services/Clinician.service';

describe('IndexClinicianComponent', () => {
  let component: IndexClinicianComponent;
  let fixture: ComponentFixture<IndexClinicianComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexClinicianComponent
      ],
      providers: [
        ClinicianService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexClinicianComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});