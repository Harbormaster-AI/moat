
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataBreachComponent } from './index.component';
import { DataBreachService } from '../../../services/DataBreach.service';

describe('IndexDataBreachComponent', () => {
  let component: IndexDataBreachComponent;
  let fixture: ComponentFixture<IndexDataBreachComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataBreachComponent
      ],
      providers: [
        DataBreachService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataBreachComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});