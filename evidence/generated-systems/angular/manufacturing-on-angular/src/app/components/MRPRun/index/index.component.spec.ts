
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMRPRunComponent } from './index.component';
import { MRPRunService } from '../../../services/MRPRun.service';

describe('IndexMRPRunComponent', () => {
  let component: IndexMRPRunComponent;
  let fixture: ComponentFixture<IndexMRPRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMRPRunComponent
      ],
      providers: [
        MRPRunService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMRPRunComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});