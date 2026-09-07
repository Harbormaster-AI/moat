
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCycleCountEntryComponent } from './index.component';
import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';

describe('IndexCycleCountEntryComponent', () => {
  let component: IndexCycleCountEntryComponent;
  let fixture: ComponentFixture<IndexCycleCountEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCycleCountEntryComponent
      ],
      providers: [
        CycleCountEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCycleCountEntryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});