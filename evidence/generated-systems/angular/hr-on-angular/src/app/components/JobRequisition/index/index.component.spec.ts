
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexJobRequisitionComponent } from './index.component';
import { JobRequisitionService } from '../../../services/JobRequisition.service';

describe('IndexJobRequisitionComponent', () => {
  let component: IndexJobRequisitionComponent;
  let fixture: ComponentFixture<IndexJobRequisitionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexJobRequisitionComponent
      ],
      providers: [
        JobRequisitionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexJobRequisitionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});