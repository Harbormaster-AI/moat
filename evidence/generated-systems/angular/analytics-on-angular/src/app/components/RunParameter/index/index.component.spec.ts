
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRunParameterComponent } from './index.component';
import { RunParameterService } from '../../../services/RunParameter.service';

describe('IndexRunParameterComponent', () => {
  let component: IndexRunParameterComponent;
  let fixture: ComponentFixture<IndexRunParameterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRunParameterComponent
      ],
      providers: [
        RunParameterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRunParameterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});