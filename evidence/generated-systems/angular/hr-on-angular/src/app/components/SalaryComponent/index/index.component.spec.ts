
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSalaryComponentComponent } from './index.component';
import { SalaryComponentService } from '../../../services/SalaryComponent.service';

describe('IndexSalaryComponentComponent', () => {
  let component: IndexSalaryComponentComponent;
  let fixture: ComponentFixture<IndexSalaryComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSalaryComponentComponent
      ],
      providers: [
        SalaryComponentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSalaryComponentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});