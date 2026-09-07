
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLeaveRequestComponent } from './index.component';
import { LeaveRequestService } from '../../../services/LeaveRequest.service';

describe('IndexLeaveRequestComponent', () => {
  let component: IndexLeaveRequestComponent;
  let fixture: ComponentFixture<IndexLeaveRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLeaveRequestComponent
      ],
      providers: [
        LeaveRequestService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLeaveRequestComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});