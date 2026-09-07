
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexShiftAssignmentComponent } from './index.component';
import { ShiftAssignmentService } from '../../../services/ShiftAssignment.service';

describe('IndexShiftAssignmentComponent', () => {
  let component: IndexShiftAssignmentComponent;
  let fixture: ComponentFixture<IndexShiftAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexShiftAssignmentComponent
      ],
      providers: [
        ShiftAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexShiftAssignmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});