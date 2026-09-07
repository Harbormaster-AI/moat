
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPerformanceCycleComponent } from './index.component';
import { PerformanceCycleService } from '../../../services/PerformanceCycle.service';

describe('IndexPerformanceCycleComponent', () => {
  let component: IndexPerformanceCycleComponent;
  let fixture: ComponentFixture<IndexPerformanceCycleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPerformanceCycleComponent
      ],
      providers: [
        PerformanceCycleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPerformanceCycleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});