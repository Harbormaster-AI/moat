
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexComplianceProgramComponent } from './index.component';
import { ComplianceProgramService } from '../../../services/ComplianceProgram.service';

describe('IndexComplianceProgramComponent', () => {
  let component: IndexComplianceProgramComponent;
  let fixture: ComponentFixture<IndexComplianceProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexComplianceProgramComponent
      ],
      providers: [
        ComplianceProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexComplianceProgramComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});