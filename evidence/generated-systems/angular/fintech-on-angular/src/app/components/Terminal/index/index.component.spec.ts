
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTerminalComponent } from './index.component';
import { TerminalService } from '../../../services/Terminal.service';

describe('IndexTerminalComponent', () => {
  let component: IndexTerminalComponent;
  let fixture: ComponentFixture<IndexTerminalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTerminalComponent
      ],
      providers: [
        TerminalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTerminalComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});