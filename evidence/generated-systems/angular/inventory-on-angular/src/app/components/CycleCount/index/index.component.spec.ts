
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCycleCountComponent } from './index.component';
import { CycleCountService } from '../../../services/CycleCount.service';

describe('IndexCycleCountComponent', () => {
  let component: IndexCycleCountComponent;
  let fixture: ComponentFixture<IndexCycleCountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCycleCountComponent
      ],
      providers: [
        CycleCountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCycleCountComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});