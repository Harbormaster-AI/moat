
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTimeEntryComponent } from './index.component';
import { TimeEntryService } from '../../../services/TimeEntry.service';

describe('IndexTimeEntryComponent', () => {
  let component: IndexTimeEntryComponent;
  let fixture: ComponentFixture<IndexTimeEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTimeEntryComponent
      ],
      providers: [
        TimeEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTimeEntryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});